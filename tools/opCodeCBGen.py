

RegMap = ['B','C','D','E','H','L','(HL)','A']
FuncMap = [
        'RLC', 'RRC', 
        'RL', 'RR', 
        'SLA', 'SRA',
        'SWAP', 'SRL',
        'BIT', 'BIT', 
        'BIT', 'BIT', 
        'BIT', 'BIT', 
        'BIT', 'BIT',
        'RES', 'RES',
        'RES', 'RES',
        'RES', 'RES',
        'RES', 'RES',
        'SET', 'SET',
        'SET', 'SET',
        'SET', 'SET',
        'SET', 'SET']

DefTmpl = "\t0x{opCode}: {{\n\t\tFunc: (*Cpu).opCodeCB{opCode},\n\t\tCycles: nil,\n\t\tLength: 2,\n\t\tMnemonic: \"{funcName} {bitSnip}{regName}\",\n\t}},\n"

FuncTmps = {
        'RLC': '''

// {funcName} {regName}
func (cpu *Cpu)opCodeCB{opCode}() byte {{
	{getCode}
	data = data << 1 | data >> 7
	{setCode}
	cpu.resetFlagZNHC()
	if data & 0x80 == 0x80 {{
		cpu.setFlagC()
	}}
	return {ret}
}}
        ''',
        'RRC': '''

// {funcName} {regName}
func (cpu *Cpu)opCodeCB{opCode}() byte {{
	{getCode}
	carry := data & 0x01 == 0x01
	data = (data >> 1) | ((data & 0x01) << 7)
	{setCode}
	cpu.resetFlagZNHC()
	if data == 0 {{
		cpu.setFlagZ()
	}}
	if carry {{
		cpu.setFlagC()
	}}
	return {ret}
}}

        ''',
        'RL': '''

// {funcName} {regName}
func (cpu *Cpu)opCodeCB{opCode}() byte {{
	{getCode}
	carry := data & 0x80 == 0x80
	data = data << 1
	if cpu.getFlagC() {{
		data |= 0x01
	}}
	{setCode}
	cpu.resetFlagZNHC()
	if data == 0 {{
		cpu.setFlagZ()
	}}
	if carry {{
		cpu.setFlagC()
	}}
	return {ret}
}}
        ''',
        'RR': '''

// {funcName} {regName}
func (cpu *Cpu)opCodeCB{opCode}() byte {{
	{getCode}
	carry := data & 0x01 == 0x01
	data = data >> 1
	if cpu.getFlagC() {{
		data |= 0x80
	}}
	{setCode}
	cpu.resetFlagZNHC()
	if data == 0 {{
		cpu.setFlagZ()
	}}
	if carry {{
		cpu.setFlagC()
	}}
	return {ret}
}}
        ''',
        'SLA': '''

// {funcName} {regName}
func (cpu *Cpu)opCodeCB{opCode}() byte {{
	{getCode}
	carry := data & 0x80 == 0x80
	data = data << 1
	{setCode}
	cpu.resetFlagZNHC()
	if data == 0 {{
		cpu.setFlagZ()
	}}
	if carry {{
		cpu.setFlagC()
	}}
	return {ret}
}}
        ''',
        'SRA': '''

// {funcName} {regName}
func (cpu *Cpu)opCodeCB{opCode}() byte {{
	{getCode}
	carry := data & 0x01 == 0x01
	data = (data >> 1) | (data & 0x80)
	{setCode}
	cpu.resetFlagZNHC()
	if data == 0 {{
		cpu.setFlagZ()
	}}
	if carry {{
		cpu.setFlagC()
	}}
	return {ret}
}}
        ''',
        'SWAP': '''

// {funcName} {regName}
func (cpu *Cpu)opCodeCB{opCode}() byte {{
	{getCode}
	data = data << 4 | data >> 4
	{setCode}
	cpu.resetFlagZNHC()
	if data == 0 {{
		cpu.setFlagZ()
	}}
	return {ret}
}}
        ''',
        'SRL': '''

// {funcName} {regName}
func (cpu *Cpu)opCodeCB{opCode}() byte {{
	{getCode}
	carry := data & 0x01 == 0x01
	data = data >> 1
	{setCode}
	cpu.resetFlagZNHC()
	if data == 0 {{
		cpu.setFlagZ()
	}}
	if carry {{
		cpu.setFlagC()
	}}
	return {ret}
}}
        ''',
        'BIT': '''

// {funcName} {bitSnip}{regName}
func (cpu *Cpu)opCodeCB{opCode}() byte {{
	{getCode}
	data = data & (0x1 << {bitN})
	if data == 0 {{
		cpu.setFlagZ()
	}}
	cpu.resetFlagN()
	cpu.setFlagH()
	return {ret}
}}
        ''',
        'RES': '''

// {funcName} {bitSnip}{regName}
func (cpu *Cpu)opCodeCB{opCode}() byte {{
	{getCode}
	data = data &^ (0x1 << {bitN})
	{setCode}
	if data == 0 {{
		cpu.setFlagZ()
	}}
	cpu.resetFlagN()
	cpu.setFlagH()
	return {ret}
}}
        ''',
        'SET': '''

// {funcName} {bitSnip}{regName}
func (cpu *Cpu)opCodeCB{opCode}() byte {{
	{getCode}
	data = data | (0x1 << {bitN})
	{setCode}
	if data == 0 {{
		cpu.setFlagZ()
	}}
	cpu.resetFlagN()
	cpu.setFlagH()
	return {ret}
}}
        '''
        }


def main():
    defCode = ""
    funcCode = ""

    for i in range(256):
        arg = {}
        arg['opCode'] = '%02X' % i
        arg['regName'] = RegMap[i & 0b111]
        bitN = (i & 0b00111000) >> 3
        arg['bitN'] = bitN
        arg['funcName'] = FuncMap[(i & 0b11111000) >> 3]
    
        arg['bitSnip'] = ''
        if i & 0b11000000 > 0:
            arg['bitSnip'] = '%d,' % bitN

        defCode += DefTmpl.format(**arg)
        
        getCode = 'data := cpu.reg.{regName}'.format(**arg)
        setCode = 'cpu.reg.{regName} = data'.format(**arg)
        ret = 2

        if arg['regName'] == '(HL)':
            getCode = 'data := cpu.gb.Mem.Read(cpu.reg.getHL())'
            setCode = 'cpu.gb.Mem.Write(cpu.reg.getHL(), data)'
            if arg['funcName'] == 'BIT':
                ret = 1
            else :
                ret = 0
        
        arg['getCode'] = getCode
        arg['setCode'] = setCode
        arg['ret'] = ret
        
        funcCode += FuncTmps[arg['funcName']].format(**arg)

    out = "package gb\n\nvar OPCodesMapCB = [0x100]OPCode{\n" + defCode + "}\n" + funcCode

    with open('opcodescb_gen.go', 'w') as f:
        f.write(out)


if __name__ == "__main__" :
    main()
