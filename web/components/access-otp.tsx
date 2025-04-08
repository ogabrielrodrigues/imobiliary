import {
  InputOTP,
  InputOTPGroup,
  InputOTPSlot,
} from "@/components/ui/input-otp"

export function AccessOTP({ id }: React.ComponentProps<"input">) {
  const pattern = "^\\d+$"

  return (
    <InputOTP
      maxLength={8}
      className="w-xs md:max-w-sm"
      id={id}
      pattern={pattern}
      autoFocus
    >
      <InputOTPGroup>
        <InputOTPSlot index={0} className="sm:size-12" />
        <InputOTPSlot index={1} className="sm:size-12" />
        <InputOTPSlot index={2} className="sm:size-12" />
        <InputOTPSlot index={3} className="sm:size-12" />
        <InputOTPSlot index={4} className="sm:size-12" />
        <InputOTPSlot index={5} className="sm:size-12" />
        <InputOTPSlot index={6} className="sm:size-12" />
        <InputOTPSlot index={7} className="sm:size-12" />
      </InputOTPGroup>
    </InputOTP>
  )
}