import { Center, ElementType, Spinner, SpinnerProps } from "@hope-ui/solid"
import { getMainColor } from "~/store"

export const CenterLoading = <C extends ElementType = "div">(
  props: SpinnerProps<C>
) => {
  return (
    <Center w="$full" h="$full">
      <Spinner color={getMainColor()} {...props} />
    </Center>
  )
}