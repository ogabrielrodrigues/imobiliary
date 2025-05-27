import type { Metadata } from "next"

import { AuthHeader } from "../_components/auth-header"
import { LoginForm } from "./_components/login-form"

export const metadata: Metadata = {
  title: "Login"
}

export default function LoginPage() {
  return (
    <div className="bg-background flex min-h-svh flex-col items-center justify-center gap-6 p-6 md:p-10">
      <div className="w-full max-w-sm">
        <div className="flex flex-col gap-6">
          <AuthHeader
            title="Bem-vindo ao Imobiliary!"
            description="Não possui acesso a plataforma?"
            url="/cadastro"
            urlText="Cadastrar"
          />
          <LoginForm />
        </div>
      </div>
    </div>
  )
}
