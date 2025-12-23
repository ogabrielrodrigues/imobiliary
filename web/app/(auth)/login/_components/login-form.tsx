'use client'

import { login } from "@/actions/mutations/auth/login"
import { PasswordInput } from "@/components/password-input"
import { ResetPasswordDialog } from "@/components/reset-password-dialog"
import { Button } from "@/components/ui/button"
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form"
import { Input } from "@/components/ui/input"
import { Spinner } from "@/components/ui/spinner"
import { cn } from "@/lib/utils"
import { zodResolver } from "@hookform/resolvers/zod"
import { IconLogin2 } from "@tabler/icons-react"
import { LoaderCircle } from "lucide-react"
import { useRouter } from "next/navigation"
import { Fragment, useState } from "react"
import { useForm } from "react-hook-form"
import { toast } from "sonner"
import { z } from "zod"

const login_schema = z.object({
  email: z.string().email("o e-mail digitado deve ser válido"),
  password: z.string().min(8, "A senha deve contem ao menos 8 caracteres")
})

export type LoginRequest = z.infer<typeof login_schema>

export function LoginForm({ className, ...props }: React.ComponentProps<"form">) {
  const [loading, setLoading] = useState(false)
  const [open, setOpen] = useState(false)

  function setDialog(open: boolean) {
    setOpen(open)
  }

  const router = useRouter()

  const form = useForm<LoginRequest>({
    resolver: zodResolver(login_schema),
    defaultValues: {
      email: "",
      password: ""
    }
  })

  async function onSubmit(values: LoginRequest) {
    setLoading(true)

    const status = await login(values)

    switch (status) {
      case 200:
        toast.success("Login realizado com sucesso!", {
          description: "Estamos lhe redirecionando a Dashboard",
          duration: 2000
        })
        setTimeout(() => {
          router.push("/dashboard")
        }, 2000)
        break
      case 401:
        toast.error("Senha inválida.", {
          description: "Verifique sua senha e tente novamente",
          duration: 4000
        })
        break
      case 404:
        toast.error("Usuário não encontrado.", {
          description: "Verifique seu email e tente novamente",
          duration: 4000
        })
        break
      default:
        toast.error("Erro ao realizar login.", {
          description: "Ocorreu um erro interno ao tentar realizar login. Caso o problema persista entre em contato",
          duration: 4000
        })
        break
    }

    setLoading(false)
  }

  return (
    <Fragment>
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
                    disabled={loading}
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
                    disabled={loading}
                    {...field}
                  />
                </FormControl>
                <FormMessage id="password-error-label" />
              </FormItem>
            )}
          />

          <Button disabled={loading} type="submit" className="w-full" id="login">
            {loading ? <Spinner/> : <IconLogin2/>}Entrar
          </Button>
        </form>
      </Form>
      <ResetPasswordDialog open={open} setOpen={setDialog} />
    </Fragment>
  )
}
