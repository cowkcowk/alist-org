import { Center, Container, HStack, useColorModeValue } from "@hope-ui/solid"

import { getSetting } from "~/store"
import { CenterLoading } from "~/components"
import { Show } from "solid-js"
import { Layout } from "./layout"

export const Header = () => {
  const logos = getSetting("logo").split("\n")
  const logo = useColorModeValue(logos[0], logos.pop())
  return (
    <Center
      class="header"
      w="$full"
    >
      hi
      <Container>
        <HStack>
          <HStack class="header-right" spacing="$2">
            <Show when={true}>
            <Layout />
            </Show>
          </HStack>
        </HStack>
      </Container>
    </Center>
  )
}