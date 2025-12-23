import type { Metadata } from "next"

import { AuthHeader } from "@/components/auth/auth-header"
import { LoginForm } from "@/components/auth/login-form"

export const metadata: Metadata = {
	title: "Login",
}

export default function LoginPage() {
	return (
		<div className="bg-primary/3 flex min-h-svh flex-col items-center justify-center gap-6 p-6 md:p-10">
			<div className="w-full max-w-sm">
				<div className="flex flex-col gap-4">
					<AuthHeader
						description="Não possui acesso a plataforma?"
						href="/cadastro"
						label="Cadastrar"
					/>

					<LoginForm />
				</div>
			</div>
		</div>
	)
}
