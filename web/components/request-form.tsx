import { Fingerprint, HousePlus, IdCard, KeyRound, Mail, Pen, UserRound } from "lucide-react"

import { cn } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import { Label } from "@/components/ui/label"
import { AccessOTP } from "./access-otp"
import { Input } from "./ui/input"
import { Textarea } from "./ui/textarea"

export function RequestForm({
  className,
  ...props
}: React.ComponentProps<"div">) {
  return (
    <div className={cn("flex flex-col gap-6", className)} {...props}>
      <form>
        <div className="flex flex-col items-center gap-6">
          <div className="flex flex-col items-center gap-2">
            <div
              className="flex flex-col items-center gap-2 font-medium"
            >
              <div className="flex size-8 items-center justify-center rounded-md">
                <HousePlus className="size-8" />
              </div>
              <span className="sr-only">Imobiliary</span>
            </div>
            <h1 className="text-2xl font-bold text-center">Solicite seu acesso</h1>
            <div className="text-center text-sm">
              Já possui acesso?{' '}
              <a href="/" className="underline underline-offset-4">
                Acessar
              </a>
            </div>
          </div>
          <div className="flex flex-col items-center gap-6 w-2xs sm:w-sm">
            <div className="grid gap-3 w-full">
              <Label htmlFor="fullname">
                <UserRound className="size-4" />
                Nome completo</Label>
              <div className="w-full flex justify-center">
                <Input id="fullname" />
              </div>
            </div>
            <div className="grid gap-3 w-full">
              <Label htmlFor="email">
                <Mail className="size-4" />
                E-mail</Label>
              <div className="w-full flex justify-center">
                <Input id="email" />
              </div>
            </div>
            <div className="grid gap-3 w-full">
              <Label htmlFor="creci">
                <IdCard className="size-4" />
                CRECI</Label>
              <div className="w-full flex justify-center">
                <Input id="creci" maxLength={7} />
              </div>
            </div>
            <div className="grid gap-3 w-full">
              <Label htmlFor="message">
                <Pen className="size-4" />
                Breve mensagem</Label>
              <div className="w-full flex justify-center">
                <Textarea id="message" className="resize-none h-32" />
              </div>
            </div>
            <Button type="submit" className="w-full">
              Solicitar
            </Button>
          </div>
        </div>
      </form >
    </div >
  )
}

/**
 * Olá prezado(a)(s), tenho um pequena administradora de imóveis em Colina/SP, e gostaria de acesso a plataforma para gerenciar com mais eficiência os imóveis os quais administro.
 */
