import { logout } from "@/actions/mutations/auth/logout";
import { getManager } from "@/actions/queries/manager/get-manager";
import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import { HousePlus, LogOut } from "lucide-react";
import Link from "next/link";

export async function HomeHeader() {
  const { manager } = await getManager()

  return (
    <header className="h-16 p-4 gap-4 w-full flex items-center justify-between border-b border-borde text-muted">
      <nav className="gap-4 flex items-center">
        <div className="flex items-center gap-2">
          <HousePlus className="size-5" />
          <span className="sr-only">Imobiliary</span>
          <h1 className="text-lg font-bold select-none">Imobiliary</h1>
        </div>

        <Separator orientation="vertical" className="!h-8" />

        <Link href="https://docs.imobiliary.com">
          <Button variant="ghost" size="sm">Docs</Button>
        </Link>
      </nav>

      {manager ? (
        <form action={logout} className="flex gap-2 items-center">
          <Link href="/dashboard" className="text-sm font-medium">{manager?.fullname.split(" ")[0]}</Link>
          <Button variant="outline" type="submit" size="sm">
            <LogOut className="size-4" />
            <p className="hidden sm:block">Sair</p>
          </Button>
        </form>
      ) : (
        <div className="flex items-center gap-2">
          <Link href="/login" id="login">
            <Button variant="secondary" size="sm">Login</Button>
          </Link>
          <Link href="/cadastro" id="sign">
            <Button size="sm">Cadastro</Button>
          </Link>
        </div>
      )}
    </header >
  )
}
