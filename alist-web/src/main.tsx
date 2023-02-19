/* @refresh reload */
import { Router } from "@solidjs/router"
import { render } from 'solid-js/web'

import { Index } from "./app"

declare global {
  interface Window {
    [key: string]: any
  }
}
render(
  () => (
    <Router>cd 
      <Index />
    </Router>
  ),
  document.getElementById("root") as HTMLElement
)