"use client"

import { zodResolver } from "@hookform/resolvers/zod"
import { IconCirclePlus } from "@tabler/icons-react"
import { useRouter } from "next/navigation"
import { type ComponentProps, useState } from "react"
import { useForm } from "react-hook-form"
import { toast } from "sonner"
import { useHookFormMask } from "use-mask-input"
import { z } from "zod"
import { sign } from "@/actions/mutations/auth/sign"
import { PasswordInput } from "@/components/password-input"
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
import { delay } from "@/lib/utils"
import { Spinner } from "../ui/spinner"

const signSchema = z
	.object({
		fullname: z
			.string()
			.min(10, "Seu nome deve contem ao menos 10 caracteres")
			.max(100, "Seu nome deve contem menos de 100 caracteres"),
		email: z.string().email("o e-mail digitado deve ser válido"),
		phone: z
			.string()
			.min(14, "O telefone deve conter ao menos 10 dígitos")
			.max(15, "O telefone deve conter no máximo 11 dígitos")
			.regex(
				/^\(\d{2}\)\s\d{4,5}-\d{4}$/,
				"O telefone deve conter apenas números",
			),
		password: z.string().min(8, "A senha deve contem ao menos 8 caracteres"),
		confirmPassword: z
			.string()
			.min(8, "A senha deve contem ao menos 8 caracteres"),
	})
	.refine(data => data.password === data.confirmPassword, {
		message: "As senhas não coincidem",
		path: ["confirmPassword"],
	})

export type SignRequest = z.infer<typeof signSchema>

export function SignForm(props: ComponentProps<"form">) {
	const navigate = useRouter()
	const [loading, setLoading] = useState(false)

	const form = useForm<SignRequest>({
		resolver: zodResolver(signSchema),
		defaultValues: {
			fullname: "",
			email: "",
			phone: "",
			password: "",
			confirmPassword: "",
		},
	})

	const registerWithMask = useHookFormMask(form.register)

	async function onSubmit(values: SignRequest) {
		setLoading(true)

		const status = await sign(values)

		switch (status) {
			case 201:
				toast.success("Usuário cadastrado com sucesso", {
					description: "Você já pode acessar sua conta.",
					duration: 3000,
				})

				setLoading(false)
				await delay(3000)

				navigate.push("/login")
				break
			case 409:
				toast.error("Usúario já cadastrado", {
					description: "Já existe um usuário com esse e-mail.",
					duration: 4000,
				})
				break
			default:
				toast.error("Erro ao cadastrar usuário", {
					description: "Verifique os dados e tente novamente.",
					duration: 4000,
				})
				break
		}

		setLoading(false)
	}

	return (
		<Form {...form}>
			<form
				className="flex flex-col space-y-4"
				onSubmit={form.handleSubmit(onSubmit)}
				{...props}
			>
				<FormField
					control={form.control}
					name="fullname"
					render={({ field }) => (
						<FormItem>
							<FormLabel className="gap-1">
								Nome Completo <span className="text-destructive">*</span>
							</FormLabel>
							<FormControl>
								<Input
									autoComplete="off"
									autoFocus
									disabled={loading}
									placeholder="Digite seu nome completo"
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
						<FormItem>
							<FormLabel className="gap-1">
								E-mail <span className="text-destructive">*</span>
							</FormLabel>
							<FormControl>
								<Input
									disabled={loading}
									placeholder="Digite seu melhor e-mail"
									{...field}
								/>
							</FormControl>
							<FormMessage />
						</FormItem>
					)}
				/>

				<FormField
					control={form.control}
					name="phone"
					render={({ field }) => (
						<FormItem>
							<FormLabel className="gap-1">
								Telefone/Celular <span className="text-destructive">*</span>
							</FormLabel>
							<FormControl>
								<Input
									autoComplete="off"
									disabled={loading}
									placeholder="Digite seu telefone/celular"
									{...field}
									{...registerWithMask(
										"phone",
										["(99) 9999-9999", "(99) 99999-9999"],
										{
											showMaskOnHover: false,
											showMaskOnFocus: false,
											required: true,
										},
									)}
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
									placeholder="Crie uma senha"
									{...field}
								/>
							</FormControl>
							<FormMessage />
						</FormItem>
					)}
				/>

				<FormField
					control={form.control}
					name="confirmPassword"
					render={({ field }) => (
						<FormItem>
							<FormLabel className="gap-1">
								Confirme sua senha <span className="text-destructive">*</span>
							</FormLabel>
							<FormControl>
								<PasswordInput
									autoComplete="off"
									disabled={loading}
									placeholder="Repita a senha criada"
									{...field}
								/>
							</FormControl>
							<FormMessage />
						</FormItem>
					)}
				/>

				<Button disabled={loading} type="submit">
					{loading ? <Spinner /> : <IconCirclePlus />}Cadastrar
				</Button>
			</form>
		</Form>
	)
}
