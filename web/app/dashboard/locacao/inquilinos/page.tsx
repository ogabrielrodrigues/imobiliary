import { Metadata } from "next"
import { TenantsHeader } from "./_components/tenants-header"

export const metadata: Metadata = {
  title: "Inquilinos",
  description: "Gerencie os inquilinos dos imóveis",
}

export default function TenantsPage() {
  return (
    <div className="container mx-auto flex flex-col space-y-10">
      <TenantsHeader />
    </div>
  )
}
