"use client"

import { zodResolver } from "@hookform/resolvers/zod"
import { IconLogin2 } from "@tabler/icons-react"
import { useRouter } from "next/navigation"
import { type ComponentProps, Fragment, useState } from "react"
import { useForm } from "react-hook-form"
import { toast } from "sonner"
import { z } from "zod"
import { login } from "@/actions/mutations/auth/login"
import { ResetPasswordDialog } from "@/components/auth/reset-password-dialog"
import { Button } from "@/components/ui/button"
import {
	Form,
	FormControl,
	FormField,
	FormItem,
	FormLabel,
	FormMessage,
} from "@/components/ui/form"
import { Input } from "@/components/ui/input"
import { PasswordInput } from "@/components/ui/password-input"
import { Spinner } from "@/components/ui/spinner"
import { delay } from "@/lib/utils"

const loginSchema = z.object({
	email: z.string().email("O e-mail digitado deve ser válido"),
	password: z.string().min(8, "A senha deve contem ao menos 8 caracteres"),
})

export type LoginRequest = z.infer<typeof loginSchema>

export function LoginForm(props: ComponentProps<"form">) {
	const navigate = useRouter()
	const [loading, setLoading] = useState(false)
	const [open, setOpen] = useState(false)

	function changeDialogState(open: boolean) {
		setOpen(open)
	}

	const form = useForm<LoginRequest>({
		resolver: zodResolver(loginSchema),
		defaultValues: {
			email: "",
			password: "",
		},
	})

	async function onSubmit(values: LoginRequest) {
		setLoading(true)

		const status = await login(values)

		switch (status) {
			case 200:
				toast.success("Login realizado com sucesso!", {
					description: "Você será redirecionando a dashboard",
					duration: 3000,
				})

				setLoading(false)
				await delay(3000)

				navigate.push("/dashboard")
				break
			case 401:
				toast.error("Senha inválida!", {
					description: "Verifique a senha e tente novamente",
					duration: 4000,
				})
				break
			case 404:
				toast.error("Usuário não encontrado!", {
					description: "Verifique seu e-mail e tente novamente",
					duration: 4000,
				})
				break
			default:
				toast.error("Erro ao realizar login.", {
					description:
						"Ocorreu um erro interno ao tentar realizar login. Caso o problema persista entre em contato",
					duration: 4000,
				})
				break
		}

		setLoading(false)
	}

	return (
		<Fragment>
			<Form {...form}>
				<form
					className="flex flex-col space-y-4"
					onSubmit={form.handleSubmit(onSubmit)}
					{...props}
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

					<FormField
						control={form.control}
						name="password"
						render={({ field }) => (
							<FormItem>
								<FormLabel className="gap-1">
									Senha <span className="text-destructive">*</span>
								</FormLabel>
								<FormControl>
									<PasswordInput
										autoComplete="off"
										disabled={loading}
										placeholder="Digite sua senha"
										{...field}
									/>
								</FormControl>
								<FormMessage />
							</FormItem>
						)}
					/>

					<Button disabled={loading} type="submit">
						{loading ? <Spinner /> : <IconLogin2 />}Entrar
					</Button>
				</form>
			</Form>

			<ResetPasswordDialog open={open} setOpen={changeDialogState} />
		</Fragment>
	)
}
