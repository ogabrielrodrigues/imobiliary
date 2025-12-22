import { auth } from "@/actions/queries/auth/auth"
import { getManager } from "@/actions/queries/manager/get-manager"
import { Card, CardContent, CardDescription, CardFooter, CardHeader } from "@/components/ui/card"
import { Metadata } from "next"
import { redirect } from "next/navigation"

export const metadata: Metadata = {
  title: "Conta",
  description: "Conta"
}

export default async function AccountPage() {
  const auth_id = await auth()
  if (!auth_id) {
    redirect("/login")
  }

  const { manager, status } = await getManager()
  if (!manager || status != 200) {
    redirect("/login")
  }

  return (
    <div className="space-y-6 w-full lg:w-2/3 mx-auto">
      <div className="flex justify-between items-center">
        <h1 className="text-3xl font-bold">Conta</h1>
      </div>
      <section className="grid sm:grid-cols-2 gap-6">
        <Card className="sm:col-span-2">
          <CardContent className="flex flex-col items-center sm:flex-row space-x-6 space-y-4">
            <div className="flex flex-col items-center sm:items-start gap-1">
              <h2 className="text-lg md:text-xl font-semibold">{manager?.fullname}</h2>
              <p className="text-xs md:text-base text-muted-foreground">{manager?.email}</p>
              <span className="inline-block bg-primary/10 text-primary text-xs font-medium px-2 py-1 rounded mt-1">Administrador</span>
            </div>
          </CardContent>
        </Card>

        <Card className="!gap-0">
          <CardHeader>
            <CardDescription>Nome completo</CardDescription>
          </CardHeader>
          <CardContent className="text-sm sm:text-base">{manager?.fullname}</CardContent>
        </Card>

        <Card className="!gap-0">
          <CardHeader>
            <CardDescription>Telefone</CardDescription>
          </CardHeader>
          <CardContent className="text-sm sm:text-base">{manager?.phone}</CardContent>
        </Card>

        <Card className="!gap-0 sm:col-span-2">
          <CardHeader>
            <CardDescription>E-mail</CardDescription>
          </CardHeader>
          <CardContent className="text-sm sm:text-base">{manager?.email}</CardContent>
          <CardFooter className="mt-1 text-xs text-muted-foreground">
            Caso deseje alterar seu e-mail entre em contato com o suporte.
          </CardFooter>
        </Card>
      </section>
    </div>
  )
}