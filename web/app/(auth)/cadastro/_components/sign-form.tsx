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

import { cn } from "@/lib/utils"

import { sign } from "@/actions/mutations/sign"
import { useRouter } from "next/navigation"
import { toast } from "sonner"
import { useHookFormMask } from 'use-mask-input'

export const sign_schema = z.object({
  fullname: z.string().min(10, "Seu nome deve contem ao menos 10 caracteres").max(100, "Seu nome deve contem menos de 100 caracteres"),
  creci_id: z.string().min(7, "Seu CRECI ter 7 digitos").regex(/^\d{5}-[FJ]$/, "Seu CRECI deve conter apenas números e o sufixo F ou J"),
  email: z.string().email("o e-mail digitado deve ser válido"),
  cellphone: z.string().min(14, "O telefone deve conter ao menos 10 dígitos").max(15, "O telefone deve conter no máximo 11 dígitos")
    .regex(/^\(\d{2}\)\s\d{4,5}-\d{4}$/, "O telefone deve conter apenas números"),
  password: z.string().min(8, "A senha deve contem ao menos 8 caracteres"),
  confirm_password: z.string().min(8, "A senha deve contem ao menos 8 caracteres")
}).refine((data) => data.password === data.confirm_password, {
  message: "As senhas não coincidem",
  path: ['confirm_password']
})

export function SignForm({ className, ...props }: React.ComponentProps<"form">) {
  const router = useRouter()

  const form = useForm<z.infer<typeof sign_schema>>({
    resolver: zodResolver(sign_schema),
    defaultValues: {
      fullname: "",
      creci_id: "",
      email: "",
      cellphone: "",
      password: "",
      confirm_password: ""
    }
  })

  const registerWithMask = useHookFormMask(form.register);

  async function onSubmit(values: z.infer<typeof sign_schema>) {
    const status = await sign(values)

    switch (status) {
      case 201:
        toast.success("Usuário cadastrado com sucesso", {
          description: "Você já pode acessar sua conta.",
          duration: 1500
        })
        setTimeout(() => {
          router.push("/login")
        }, 1500)
        break
      case 409:
        toast.error("Erro ao cadastrar usuário", {
          description: "Já existe um usuário com esse e-mail ou CRECI.",
          duration: 2000
        })
        break
      default:
        toast.error("Erro ao cadastrar usuário", {
          description: "Verifique os dados e tente novamente.",
          duration: 2000
        })
        break
    }
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
          name="creci_id"
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
                  {...registerWithMask("creci_id", ["99999-F", "99999-J"], {
                    showMaskOnHover: false,
                    showMaskOnFocus: false,
                    required: true,
                  })}
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
          name="cellphone"
          render={({ field }) => (
            <FormItem className="col-span-1 md:col-span-3">
              <FormLabel>Telefone / Celular</FormLabel>
              <FormControl>
                <Input
                  className="text-sm md:text-base"
                  placeholder="Seu telefone ou celular"
                  autoComplete="off"
                  {...field}
                  {...registerWithMask("cellphone", ['(99) 9999-9999', '(99) 99999-9999'], {
                    showMaskOnHover: false,
                    showMaskOnFocus: false,
                    required: true
                  })}
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
