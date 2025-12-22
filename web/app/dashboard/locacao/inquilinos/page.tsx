import { listTenants } from "@/actions/queries/tenant/list-tenant"
import { Metadata } from "next"
import { TenantsHeader } from "./_components/tenants-header"
import { TenantsSection } from "./_components/tenants-section"

export const metadata: Metadata = {
  title: "Inquilinos",
  description: "Gerencie os inquilinos dos imóveis",
}

export default async function TenantsPage() {
  const { status: tenant_status, tenants: foundTenants } = await listTenants()

  if (tenant_status != 200) {
    return (
      <div className="w-full flex justify-center">
        <p className="font-medium text-muted">Erro ao carregar inquilinos</p>
      </div>
    )
  }

  const tenants = !foundTenants ? [] : foundTenants

  return (
    <div className="container mx-auto flex flex-col space-y-10">
      <TenantsHeader tenants={tenants} />

      <TenantsSection tenants={tenants} />
    </div>
  )
}
