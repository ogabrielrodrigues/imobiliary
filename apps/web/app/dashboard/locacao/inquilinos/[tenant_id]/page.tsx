import { getTenant } from "@/actions/queries/tenant/get-tenant"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { ArrowLeft } from "lucide-react"
import { Metadata } from "next"
import Link from "next/link"
import { notFound } from "next/navigation"

type TenantDetailsPageParams = {
  params: Promise<{ tenant_id: string }>
}

export const metadata: Metadata = {
  title: "Perfil do Inquilino",
}

export default async function TenantDetailsPage({ params }: TenantDetailsPageParams) {
  const { tenant_id } = await params
  const { status, tenant } = await getTenant(tenant_id)

  if (!tenant || status !== 200) notFound()

  return (
    <div className="mx-auto max-w-3xl w-full flex flex-col space-y-4">
      <div className="flex justify-between gap-4">
        <div>
          <h1 className="text-2xl font-bold">{tenant.fullname}</h1>
          <span className="text-muted-foreground text-sm">{tenant.address.address}</span>
        </div>
        <Link href="/dashboard/locacao/inquilinos">
          <Button variant="outline">
            <ArrowLeft className="w-4 h-4" />
            <p className="hidden md:block">Voltar</p>
          </Button>
        </Link>
      </div>
      <Card className="backdrop-blur-2xl relative z-20 overflow-hidden">
        <CardHeader>
          <CardTitle className="text-lg font-bold">Ficha Inquilino</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="flex flex-col">
            <p className="text-muted-foreground"><strong className="text-muted">Telefone:</strong> {tenant.fullname}</p>
            <p className="text-muted-foreground"><strong className="text-muted">RG:</strong> {tenant.rg}</p>
            <p className="text-muted-foreground"><strong className="text-muted">CPF:</strong> {tenant.cpf}</p>
            <p className="text-muted-foreground"><strong className="text-muted">Profissão:</strong> {tenant.occupation}</p>
            <p className="text-muted-foreground"><strong className="text-muted">Estado civil:</strong> {tenant.marital_status}</p>
            <p className="text-muted-foreground"><strong className="text-muted">Telefone:</strong> {tenant.phone}</p>
          </div>
        </CardContent>
      </Card>
    </div>
  )
}