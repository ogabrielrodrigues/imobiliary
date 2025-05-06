


import { logout } from "@/actions/mutations/auth/logout";
import { auth } from "@/actions/queries/auth/auth";
import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import { HousePlus, LogOut } from "lucide-react";
import Link from "next/link";

export async function HomeHeader() {
  const user = await auth()

  return (
    <header className="h-16 border-b border-border w-full bg-zinc-950/95 backdrop-blur-3xl p-4 flex items-center justify-between gap-4">
      <nav className="flex gap-2 items-center">
        <div className="flex items-center gap-2">
          <HousePlus className="size-5" />
          <span className="sr-only">Imobiliary</span>
          <h1 className="text-lg font-bold hidden sm:block select-none">Imobiliary</h1>
        </div>

        <Separator orientation="vertical" />

        <Link href="https://docs.imobiliary.com">
          <Button variant="ghost" size="sm">Docs</Button>
        </Link>
      </nav>

      {user ? (<form action={logout}>
        <Button variant="outline" type="submit" size="sm">
          <LogOut className="size-4" />
          <p className="hidden sm:block">Sair</p>
        </Button>
      </form>)
        : (<div className="flex items-center gap-2">
          <Link href="/login" id="login">
            <Button variant="ghost" size="sm">Login</Button>
          </Link>
          <Link href="/cadastro" id="sign">
            <Button variant="ghost" size="sm">Cadastro</Button>
          </Link>
        </div>)}
    </header >
  )
}
