package pdf

import (
	"context"
	"sync"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func GeneratePDFFromHtml(html string, width, height float64) ([]byte, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var buf []byte
	err := chromedp.Run(ctx, convertToPDF(html, width, height, &buf))
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func convertToPDF(html string, width, height float64, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate("about:blank"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			lctx, cancel := context.WithCancel(ctx)
			defer cancel()
			var wg sync.WaitGroup
			wg.Add(1)
			chromedp.ListenTarget(lctx, func(ev interface{}) {
				if _, ok := ev.(*page.EventLoadEventFired); ok {
					cancel()
					wg.Done()
				}
			})

			frameTree, err := page.GetFrameTree().Do(ctx)
			if err != nil {
				return err
			}

			err = page.SetDocumentContent(frameTree.Frame.ID, html).Do(ctx)
			if err != nil {
				return err
			}

			wg.Wait()
			return nil
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			chromePage := page.PrintToPDF().WithPrintBackground(true)

			if width > 0.00 && height > 0.00 {
				chromePage.WithPaperWidth(width).WithPaperHeight(height)
			}

			buf, _, err := chromePage.Do(ctx)
			if err != nil {
				return err
			}

			*res = buf
			return nil
		}),
	}
}
