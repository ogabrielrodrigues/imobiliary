import { Button } from "@/components/ui/button"
import { ArrowLeft } from "lucide-react"
import { Metadata } from "next"
import Link from "next/link"
import { NewPropertyForm } from "../_components/new-property-form"

export const metadata: Metadata = {
  title: "Imóveis",
  description: "Gerencie os imóveis disponíveis para aluguel",
}

export default async function NewPropertyPage() {
  return (
    <div className="container mx-auto xl:max-w-screen-md flex flex-col space-y-4">
      <div className="flex items-center justify-between">
        <h1 className="text-2xl font-bold">Novo Imóvel</h1>
        <Link href="/dashboard/locacao/imoveis">
          <Button variant="outline">
            <ArrowLeft className="w-4 h-4" />
            <p className="hidden md:block">Voltar</p>
          </Button>
        </Link>
      </div>
      <NewPropertyForm />
    </div >
  )
}
