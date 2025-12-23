"use client"

import { zodResolver } from "@hookform/resolvers/zod"
import { IconRefresh } from "@tabler/icons-react"
import { useState } from "react"
import { useForm } from "react-hook-form"
import { toast } from "sonner"
import { z } from "zod"
import { changePassword } from "@/actions/mutations/manager/change-password"
import { Button } from "@/components/ui/button"
import {
	Dialog,
	DialogContent,
	DialogDescription,
	DialogFooter,
	DialogHeader,
	DialogTitle,
	DialogTrigger,
} from "@/components/ui/dialog"
import {
	Form,
	FormControl,
	FormField,
	FormItem,
	FormLabel,
	FormMessage,
} from "@/components/ui/form"
import { Input } from "@/components/ui/input"
import { Spinner } from "@/components/ui/spinner"

const resetPassword = z.object({
	email: z.string().email("o e-mail digitado deve ser válido"),
})

export type ResetPasswordRequest = z.infer<typeof resetPassword>

interface ResetPasswordDialogProps {
	open?: boolean
	setOpen(open: boolean): void
}

export function ResetPasswordDialog({
	open,
	setOpen,
}: ResetPasswordDialogProps) {
	const [loading, setLoading] = useState(false)

	const form = useForm<ResetPasswordRequest>({
		resolver: zodResolver(resetPassword),
		defaultValues: {
			email: "",
		},
	})

	async function onSubmit(values: ResetPasswordRequest) {
		setLoading(true)

		const status = await changePassword(values)

		switch (status) {
			case 200:
				toast.success("E-mail de redefinição enviado ao endereço fornecido!", {
					description:
						"Acesse o e-mail e siga os passos para redefinição de sua senha.",
					duration: 3000,
				})
				break
			case 404:
				toast.error("Usuário não encontrado!", {
					description: "Verifique seu email e tente novamente",
					duration: 4000,
				})
				break
			default:
				toast.error("Erro ao redefinir senha!", {
					description:
						"Ocorreu um erro interno ao tentar redefinir senha. Caso o problema persista entre em contato",
					duration: 4000,
				})
				break
		}

		setOpen(false)
		setLoading(false)
	}

	return (
		<Dialog onOpenChange={setOpen} open={open}>
			<DialogTrigger
				className="p-0 m-0 text-sm self-end text-primary underline underline-offset-2 cursor-pointer"
				onClick={() => setOpen(true)}
			>
				Esqueceu sua senha?
			</DialogTrigger>

			<DialogContent>
				<DialogHeader>
					<DialogTitle>Redefinir sua senha</DialogTitle>

					<DialogDescription>
						Insira seu e-mail para receber o link de redefinição de senha
					</DialogDescription>
				</DialogHeader>

				<Form {...form}>
					<form
						className="flex flex-col space-y-4"
						onSubmit={form.handleSubmit(onSubmit)}
					>
						<FormField
							control={form.control}
							name="email"
							render={({ field }) => (
								<FormItem>
									<FormLabel className="gap-1">
										E-mail <span className="text-destructive">*</span>
									</FormLabel>
									<FormControl>
										<Input
											autoFocus
											disabled={loading}
											placeholder="Digite seu e-mail"
											{...field}
										/>
									</FormControl>
									<FormMessage />
								</FormItem>
							)}
						/>

						<DialogFooter>
							<Button className="group" disabled={loading} type="submit">
								{loading ? (
									<Spinner />
								) : (
									<IconRefresh className="group-hover:rotate-90 transition-transform" />
								)}{" "}
								Redefinir senha
							</Button>
						</DialogFooter>
					</form>
				</Form>
			</DialogContent>
		</Dialog>
	)
}
