import type { Metadata } from "next"

import { AuthHeader } from "@/components/auth/auth-header"
import { SignForm } from "@/components/auth/sign-form"

export const metadata: Metadata = {
	title: "Cadastrar",
}

export default function SignPage() {
	return (
		<div className="bg-primary/3 flex min-h-svh flex-col items-center justify-center gap-6 p-6 md:p-10">
			<div className="w-full max-w-sm">
				<div className="flex flex-col gap-4">
					<AuthHeader
						description="Já possui acesso a plataforma?"
						href="/login"
						label="Acessar"
					/>

					<SignForm />
				</div>
			</div>
		</div>
	)
}
