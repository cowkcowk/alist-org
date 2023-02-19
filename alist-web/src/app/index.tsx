import { HopeProvider, NotificationsProvider } from "@hope-ui/solid"
import { I18nContext } from "@solid-primitives/i18n"
import { ErrorBoundary } from "solid-js"
import { Error } from "~/components"

import App from "./App"
import { i18n } from "./i18n"
import { theme } from "./theme"

const Index = () => {
  return (
    <HopeProvider config={theme}>
      <ErrorBoundary
        fallback={(err) => {
          console.error("error", err)
          return <Error msg={`System error: ${err}`} h="100vh"/>
        }}
      >
        <I18nContext.Provider value={i18n}>

        </I18nContext.Provider>
      </ErrorBoundary>
      <App />
    </HopeProvider>
  )
}

export { Index }
