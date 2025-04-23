'use client'

import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import { z } from "zod"

import { PasswordInput } from "@/components/password-input"
import { Button } from "@/components/ui/button"
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage
} from "@/components/ui/form"
import { Input } from "@/components/ui/input"

import { login } from "@/actions/queries/login"
import { cn } from "@/lib/utils"
import { useRouter } from "next/navigation"
import { toast } from "sonner"

export const login_schema = z.object({
  email: z.string().email("o e-mail digitado deve ser válido"),
  password: z.string().min(8, "A senha deve contem ao menos 8 caracteres")
})

export function LoginForm({ className, ...props }: React.ComponentProps<"form">) {
  const router = useRouter()

  const form = useForm<z.infer<typeof login_schema>>({
    resolver: zodResolver(login_schema),
    defaultValues: {
      email: "",
      password: ""
    }
  })

  async function onSubmit(values: z.infer<typeof login_schema>) {
    const status = await login(values)

    switch (status) {
      case 200:
        toast.success("Login realizado com sucesso!", { duration: 1500 })
        setTimeout(() => {
          router.push("/dashboard")
        }, 1500)
        break
      case 401:
        toast.error("E-mail ou senha inválidos.", { duration: 1500 })
        break
      case 404:
        toast.error("Usuário não encontrado.", { duration: 1500 })
        break
      default:
        toast.error("Erro ao realizar login.", { duration: 1500 })
        break
    }
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
              <FormMessage id="email-error-label" />
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
              <FormMessage id="password-error-label" />
            </FormItem>
          )}
        />
        <Button type="submit" className="w-full" id="login">
          Entrar
        </Button>
      </form>
    </Form>
  )
}
