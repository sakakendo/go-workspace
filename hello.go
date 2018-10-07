package __main__
import πg "grumpy"
var Code *πg.Code

func init() {
	Code = πg.NewCode("<module>", "hello.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßmain := πg.InternStr("main")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []πg.Param
		_ = πTemp002
		var πTemp003 *πg.Object
		_ = πTemp003
		var πTemp004 *πg.Object
		_ = πTemp004
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 1: def main():
			πF.SetLineno(1)
			πTemp002 = make([]πg.Param, 0)
			πTemp001 = πg.NewFunction(πg.NewCode("main", "hello.py", πTemp002, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var πTemp001 []*πg.Object
				_ = πTemp001
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					// line 2: print("hello world")
					πF.SetLineno(2)
					πTemp001 = make([]*πg.Object, 1)
					πTemp001[0] = πg.NewStr("hello world").ToObject()
					if πE = πg.Print(πF, πTemp001, true); πE != nil {
						continue
					}
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßmain.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 5: main()
			πF.SetLineno(5)
			if πTemp003, πE = πg.ResolveGlobal(πF, ßmain); πE != nil {
				continue
			}
			if πTemp004, πE = πTemp003.Call(πF, nil, nil); πE != nil {
				continue
			}
		}
		return nil, πE
	})
	πg.RegisterModule("__main__", Code)
}

func main(){
    init()
}
