'use client'

import { useForm } from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"
import { z } from "zod"

import { PasswordInput } from "@/components/password-input"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage
} from "../ui/form"

import { cn } from "@/lib/utils"

const sign_schema = z.object({
  fullname: z.string().min(10, "Seu nome deve contem ao menos 10 caracteres").max(100, "Seu nome deve contem menos de 100 caracteres"),
  creci: z.string().min(7, "Seu CRECI deve ser válido"),
  email: z.string().email("o e-mail digitado deve ser válido"),
  password: z.string().min(8, "A senha deve contem ao menos 8 caracteres"),
  confirm_password: z.string().min(8, "A senha deve contem ao menos 8 caracteres")
}).refine((data) => data.password === data.confirm_password, {
  message: "As senhas não coincidem",
  path: ['confirm_password']
})

export function SignForm({ className, ...props }: React.ComponentProps<"form">) {
  const form = useForm<z.infer<typeof sign_schema>>({
    resolver: zodResolver(sign_schema),
    defaultValues: {
      fullname: "",
      creci: "",
      email: "",
      password: "",
      confirm_password: ""
    }
  })

  function onSubmit(values: z.infer<typeof sign_schema>) {
    // TODO: implementar lógica de login de usário
    console.log(values)
  }

  return (
    <Form {...form}>
      <form
        className={cn("grid grid-cols-1 md:!grid-cols-3 space-y-4 gap-x-2", className)}
        {...props}
        onSubmit={form.handleSubmit(onSubmit)}
      >
        <FormField
          control={form.control}
          name="fullname"
          render={({ field }) => (
            <FormItem className="col-span-1 md:col-span-2">
              <FormLabel>Nome Completo</FormLabel>
              <FormControl>
                <Input
                  className="text-sm md:text-base"
                  placeholder="Seu nome completo"
                  autoComplete="off"
                  autoFocus
                  {...field}
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="creci"
          render={({ field }) => (
            <FormItem className="col-span-1">
              <FormLabel>Seu CRECI</FormLabel>
              <FormControl>
                <Input
                  className="text-sm md:text-base"
                  placeholder="Seu CRECI"
                  maxLength={7}
                  autoComplete="off"
                  {...field}
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="email"
          render={({ field }) => (
            <FormItem className="col-span-1 md:col-span-3">
              <FormLabel>E-mail</FormLabel>
              <FormControl>
                <Input
                  className="text-sm md:text-base"
                  placeholder="Seu e-mail de acesso"
                  autoComplete="off"
                  {...field}
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="password"
          render={({ field }) => (
            <FormItem className="col-span-1 md:col-span-3">
              <FormLabel>Senha</FormLabel>
              <FormControl>
                <PasswordInput
                  className="text-sm md:text-base"
                  placeholder="Sua senha de acesso"
                  autoComplete="off"
                  {...field}
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="confirm_password"
          render={({ field }) => (
            <FormItem className="col-span-1 md:col-span-3">
              <FormLabel>Confirme sua senha</FormLabel>
              <FormControl>
                <PasswordInput
                  className="text-sm md:text-base"
                  placeholder="Confirme sua senha de acesso"
                  autoComplete="off"
                  {...field}
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <Button type="submit" className="w-full col-span-1 md:col-span-3">
          Cadastrar
        </Button>
      </form>
    </Form>
  )
}
