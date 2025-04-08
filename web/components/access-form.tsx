import { HousePlus, KeyRound } from "lucide-react"

import { cn } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import { Label } from "@/components/ui/label"
import { AccessOTP } from "./access-otp"

export function AccessForm({
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
            <h1 className="text-2xl font-bold text-center">Bem-vindo ao Imobiliary!</h1>
            <div className="text-center text-sm">
              Ainda não possui acesso?{' '}
              <a href="/solicitar" className="underline underline-offset-4">
                Solicitar acesso
              </a>
            </div>
          </div>
          <div className="flex flex-col items-center gap-6 w-2xs sm:w-sm">
            <div className="grid gap-3 w-sm">
              <Label htmlFor="access">
                <KeyRound className="size-4" />
                Código de acesso</Label>
              <div className="max-w-sm flex justify-center">
                <AccessOTP id="access" />
              </div>
            </div>
            <Button type="submit" className="w-full">
              Acessar
            </Button>
          </div>
        </div>
      </form>
    </div>
  )
}
