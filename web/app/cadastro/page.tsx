import type { Metadata } from "next"
import Link from "next/link"

import { HousePlus } from "lucide-react"

import { SignForm } from "@/components/forms/sign-form"
import { Button } from "@/components/ui/button"

export const metadata: Metadata = {
  title: "Cadastrar"
}

export default function SignPage() {
  return (
    <div className="bg-background flex min-h-svh flex-col items-center justify-center gap-6 p-6 md:p-10">
      <div className="w-full max-w-sm">
        <div className="flex flex-col gap-6">
          <div className="flex flex-col gap-6">
            <div className="flex flex-col items-center gap-2">
              <div className="flex flex-col items-center gap-2 font-medium">
                <Link href="/" className="flex size-8 items-center justify-center rounded-md">
                  <Button variant="ghost" size="icon">
                    <HousePlus className="size-8" />
                  </Button>
                </Link>
                <span className="sr-only">Imobiliary</span>
              </div>
              <h1 className="text-xl font-bold text-center">Bem-vindo ao Imobiliary!</h1>
              <div className="text-center text-sm">
                Já possui acesso a plataforma?{" "}
                <a href="/login" className="underline underline-offset-4">
                  Acessar
                </a>
              </div>
            </div>
          </div>
          <SignForm />
        </div>
      </div>
    </div>
  )
}