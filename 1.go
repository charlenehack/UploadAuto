func main() {
	options := []chromedp.ExecAllocatorOption{
		//	chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
		chromedp.NoDefaultBrowserCheck,
		//	chromedp.NoFirstRun,
	}

	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), options...)
	defer cancel()

	ctx, _ := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf))
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 15 * time.Second)
	defer cancel()

	chromedp.Run(ctx, chromedp.Navigate("https://emkei.cz/"))

	for {
		err := chromedp.Run(ctx,
			//chromedp.Navigate("https://emkei.cz/"),
			//chromedp.Sleep(15*time.Second),
			//chromedp.WaitVisible(`#sendfrm`, chromedp.ByID),
			chromedp.SendKeys(`input[name=fromname]`, fromname),
			chromedp.Sleep(10*time.Second),
			chromedp.SendKeys(`input[name=from]`, from),
			chromedp.Sleep(10*time.Second),
			chromedp.SendKeys(`input[name=rcpt]`, tos),
			chromedp.Sleep(10*time.Second),
			chromedp.SendKeys(`input[name=subject]`, subject),
			chromedp.Sleep(10*time.Second),
			chromedp.SendKeys(`textarea[name=text]`, msg),
			chromedp.Sleep(10*time.Second),
			chromedp.Click(".btn.sbold.slarger", chromedp.ByQuery),
			//chromedp.Click(`input[type=reset]`, chromedp.ByQuery),
			chromedp.Sleep(2*time.Second),

			chromedp.NodeIDs("document", &ids, chromedp.ByJSPath),
			chromedp.ActionFunc(func(ctx context.Context) error {
				var err error
				body, err = dom.GetOuterHTML().WithNodeID(ids[0]).Do(ctx)
				return err
			}),
		)

		time.Sleep(5 * time.Second)
		if err != nil {
			log.Fatal(err)
		}
		if strings.Contains(body, "successfully") {
			fmt.Println("send successed.")
			WriteLog(tos, "sucess")
		} else {
			fmt.Println("send failed.")
			WriteLog(tos, "fail")
			chromedp.ListenTarget(ctx, func(ev interface{}) {
				if ev, ok := ev.(*page.EventJavascriptDialogOpening); ok {
					fmt.Println("closing alert:", ev.Message)
					go func() {
						//自动关闭alert对话框
						if err := chromedp.Run(ctx,
							//注释掉下一行可以更清楚地看到效果
							page.HandleJavaScriptDialog(true),
						); err != nil {
							fmt.Println(err)
						}
					}()
				}
			})
			chromedp.Run(ctx,
				chromedp.Click(`input[type=reset]`, chromedp.ByQuery),
				chromedp.Sleep(5*time.Second),
			)

		}
		time.Sleep(10 * time.Second)
	}
}
