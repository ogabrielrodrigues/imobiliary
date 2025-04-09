import type { Metadata } from "next"
import Link from "next/link"

import { HousePlus } from "lucide-react"

import { TwoFactorForm } from "@/components/forms/2fa-form"

export const metadata: Metadata = {
  title: "2FA"
}

function formatEmail(email: string): string {
  const [address, domain] = email.split("@")

  let formated_address = address.slice(0, 4)
  formated_address += "*".repeat(address.length / 2)

  return [formated_address, "@", domain].join("")
}

export default function TwoFactorAuthPage() {
  return <div className="bg-background flex min-h-svh flex-col items-center justify-center gap-6 p-6 md:p-10">
    <div className="flex flex-col gap-1 w-60 sm:w-2xs text-center">
      <h1 className="text-2xl font-bold">Confirme o código enviado:</h1>
      <span className="text-muted-foreground text-sm">{formatEmail("gabrielrodrigues@outlook.com")}</span>
    </div>
    <div className="w-max">
      <TwoFactorForm />
    </div>
  </div>
}


