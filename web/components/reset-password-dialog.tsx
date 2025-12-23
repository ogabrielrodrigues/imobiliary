"use client"

import { cn } from '@/lib/utils'
import { zodResolver } from '@hookform/resolvers/zod'
import { useState } from 'react'
import { useForm } from 'react-hook-form'
import { toast } from 'sonner'
import { z } from 'zod'
import { Button } from './ui/button'
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle, DialogTrigger } from "./ui/dialog"
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from './ui/form'
import { Input } from './ui/input'
import { changePassword } from '@/actions/mutations/manager/change-password'

const reset_password = z.object({
  email: z.string().email("o e-mail digitado deve ser válido"),
})

export type ResetPasswordRequest = z.infer<typeof reset_password>

type ResetPasswordDialogProps = {
  open?: boolean
  setOpen(open: boolean): void
}

export function ResetPasswordDialog({ open, setOpen }: ResetPasswordDialogProps) {
  const [loading, setLoading] = useState(false)

  const form = useForm<ResetPasswordRequest>({
    resolver: zodResolver(reset_password),
    defaultValues: {
      email: "",
    }
  })

  async function onSubmit(values: ResetPasswordRequest) {
    setLoading(true)

    const status = await changePassword(values)

    switch (status) {
      case 200:
        toast.success("E-mail de redefinição enviado ao endereço fornecido.", {
          description: "Acesse o e-mail e siga os passos para redefinição de sua senha.",
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
        toast.error("Erro ao redefinir senha.", {
          description: "Ocorreu um erro interno ao tentar redefinir senha. Caso o problema persista entre em contato",
          duration: 4000
        })
        break
    }

    setOpen(false)
    setLoading(false)
  }

  return (
    <Dialog open={open}>
      <DialogTrigger className="p-0 m-0 self-end text-primary underline cursor-pointer" onClick={() => setOpen(true)}>Esqueceu sua senha?</DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Redefinir sua senha</DialogTitle>
          <DialogDescription>Insira seu e-mail para receber o link de redefinição de senha</DialogDescription>
        </DialogHeader>
        <Form {...form}>
          <form
            className={cn("flex flex-col space-y-4")}
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
            <DialogFooter>
              <Button type='submit'>Redefinir senha</Button>
            </DialogFooter>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  )
}
