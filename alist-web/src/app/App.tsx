import { Progress, ProgressIndicator } from "@hope-ui/solid"
import { Route, Routes, useIsRouting } from "@solidjs/router"
import {
  Component,
  createSignal,
  lazy,
  Match,
  onCleanup,
  Switch,
} from "solid-js"
import { Portal } from "solid-js/web"
import { base_path } from "~/utils"

const Home = lazy(() => import("~/pages/home/Layout"))
const About = lazy(() => import("~/pages/manage/About"))

const App: Component = () => {
  return (
    <>
      <Portal>
        <Progress
          indeterminate
          size="xs"
          position="fixed"
          top="0"
          left="0"
          right="0"
          zIndex="$banner"
          d="block"
        >
          <ProgressIndicator />
        </Progress>
      </Portal>
      <Switch
        fallback={
          <Routes base={base_path}>
            <Route
              path="*"
              element={
                <Home />
              }/>
          </Routes>
        }>

      </Switch>
    </>
  )
}

export default App
