import { Button } from "@/components/ui/button"
import { ArrowLeft } from "lucide-react"
import { Metadata } from "next"
import Link from "next/link"
import { NewOwnerForm } from "../_components/new-owner-form"

export const metadata: Metadata = {
  title: "Proprietários - Novo",
}

export default async function NewOwnersPage() {
  return (
    <div className="mx-auto xl:max-w-xl flex flex-col space-y-4">
      <div className="flex items-center justify-between">
        <h1 className="text-2xl font-bold">Novo Proprietário</h1>
        <Link href="/dashboard/locacao/proprietarios">
          <Button variant="outline">
            <ArrowLeft className="w-4 h-4" />
            <p className="hidden md:block">Voltar</p>
          </Button>
        </Link>
      </div>
      <NewOwnerForm />
    </div >
  )
}
