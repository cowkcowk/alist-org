import { HopeProvider, NotificationsProvider } from "@hope-ui/solid"

import App from "./App"

const Index = () => {
  return (
    <HopeProvider>
      <App />
    </HopeProvider>
  )
}

export { Index }
