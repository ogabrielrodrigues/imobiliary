'use client'

import { useForm } from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"
import { z } from "zod"

import { Button } from "@/components/ui/button"

import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage
} from "@/components/ui/form"

import { cn } from "@/lib/utils"

import {
  InputOTP,
  InputOTPGroup,
  InputOTPSlot,
} from "@/components/ui/input-otp";


const tfa_schema = z.object({
  code: z.string().min(6, "o codigo digitado tem de ter 6 digitos"),
})

export function TwoFactorForm({ className, ...props }: React.ComponentProps<"form">) {
  const form = useForm<z.infer<typeof tfa_schema>>({
    resolver: zodResolver(tfa_schema),
    defaultValues: {
      code: "",
    }
  })

  function onSubmit(values: z.infer<typeof tfa_schema>) {
    // TODO: implementar lógica de criação de usuário
    console.log(values)
  }

  return (
    <Form {...form}>
      <form
        className={cn("flex flex-col space-y-4", className)}
        {...props}
        onSubmit={form.handleSubmit(onSubmit)}
      >
        <FormField
          control={form.control}
          name="code"
          render={({ field }) => (
            <FormItem className="flex flex-col w-full">
              <FormLabel>Código de confirmação</FormLabel>
              <FormControl>
                <InputOTP maxLength={6} {...field} autoFocus>
                  <InputOTPGroup>
                    <InputOTPSlot index={0} className="size-10 sm:size-12" />
                    <InputOTPSlot index={1} className="size-10 sm:size-12" />
                    <InputOTPSlot index={2} className="size-10 sm:size-12" />
                    <InputOTPSlot index={3} className="size-10 sm:size-12" />
                    <InputOTPSlot index={4} className="size-10 sm:size-12" />
                    <InputOTPSlot index={5} className="size-10 sm:size-12" />
                  </InputOTPGroup>
                </InputOTP>
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <Button type="submit" className="w-60 sm:w-2xs">
          Confirmar
        </Button>
      </form>
    </Form>
  )
}
