'use client'

import { useForm } from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"
import { z } from "zod"
import { toast } from "sonner"

import { AlertCircle, Mail } from "lucide-react"

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
} from "@/components/ui/form"

import { cn } from "@/lib/utils"
import { User } from "@/types/user"
import { redirect, useRouter } from "next/navigation"

const auth_schema = z.object({
  email: z.string().email("o e-mail digitado deve ser válido"),
  password: z.string().min(8, "A senha deve contem ao menos 8 caracteres")
})

export function AuthForm({ className, ...props }: React.ComponentProps<"form">) {
  const navigate = useRouter()

  const form = useForm<z.infer<typeof auth_schema>>({
    resolver: zodResolver(auth_schema),
    defaultValues: {
      email: "",
      password: ""
    }
  })

  async function onSubmit(values: z.infer<typeof auth_schema>) {
    const response = await fetch("/api/mock/login", {
      method: "POST",
      body: JSON.stringify(values)
    })

    if (response.status === 400) {
      return toast("E-mail ou senha inválidos.", {
        icon: <AlertCircle className="size-4" />,
      })
    }

    return redirect("/dashboard")
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
          name="email"
          render={({ field }) => (
            <FormItem>
              <FormLabel>E-mail</FormLabel>
              <FormControl>
                <Input
                  className="text-sm md:text-base"
                  placeholder="Seu e-mail de acesso"
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
          name="password"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Senha</FormLabel>
              <FormControl>
                <PasswordInput
                  className="text-sm md:text-base"
                  type="password"
                  placeholder="Sua senha de acesso"
                  autoComplete="off"
                  {...field}
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <Button type="submit" className="w-full">
          Entrar
        </Button>
      </form>
    </Form>
  )
}
