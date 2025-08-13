func MergeContext(ctxs ...context.Context)(context.Context, context.CancelFunc){
	ctx, cancel := context.WithCancel(context.Background())

	go func(){
		defer cancel()
		cases := make([]reflect.SelectCase, len(ctxs))
		for i, c := range ctxs {
			cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(c.Done())}
		}
		
		reflect.Select(cases)
	}()

	return ctx, cancel
}
