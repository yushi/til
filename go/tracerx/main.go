package main

import (
	"time"

	"github.com/rubyist/tracerx"
)

func main() {
	// watch する変数名を設定する
	tracerx.DefaultKey = "MY"

	// MY_TRACE 環境変数が true or 1 or 2 の場合に STDERR にメッセージ出力
	// 変数がフルパスの (IsAbs が true を返す) 場合はそのファイルにメッセージを書き込む
	tracerx.Printf("MY_TRACE is enabled")

	t := time.Now()
	// MY_TRACE 環境変数が有効 かつ MY_TRACE_PERFORMANCE が true or 1 の場合
	// 出力先は MY_TRACE に依存
	tracerx.PerformanceSince("MY_TRACE_PERFORMANCE", t)

	// 以下 (Key suffix な関数) の場合 MY_TRACE ではなく任意の変数名 (この場合 BAR) を参照する
	tracerx.PrintfKey("BAR", "BAR_TRACE is enabled")
	tracerx.PerformanceSinceKey("BAR", "command x", t)
}
