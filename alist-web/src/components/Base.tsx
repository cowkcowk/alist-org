import { Box, Center, Flex, Heading, useColorModeValue } from "@hope-ui/solid"
import { mergeProps, Show } from "solid-js"
import { SwitchColorMode } from "./SwitchColorMode"

export const Error = (props: {
  msg: string
  disableColor?: boolean
  h?: string
}) => {
  const merged = mergeProps(
    {
      h: "$full",
    },
    props
  )
  console.log(merged.h)
  return (
    <Center h={merged.h} p="$2" flexDirection="column">
      <Box
        rounded="$lg"
        px="$4"
        py="$6"
        bgColor={useColorModeValue("white", "$neutral3")()}
      >
        <Heading
          css={{
            wordBreak: "break-all",
          }}
        >
          {props.msg}
        </Heading>
        <Show when={!props.disableColor}>
          <Flex mt="$2" justifyContent="end">
            <SwitchColorMode />
          </Flex>
        </Show>
      </Box>
    </Center>
  )
}