import { Button } from "@/components/ui/button";
import Link from "next/link";

export default function NotFound() {
  return <div className="w-full h-full flex flex-col items-center justify-center gap-4">
    <h1 className="text-7xl font-bold">404</h1>

    <p className="text-2xl font-bold text-center">Inquilino não encontrado</p>

    <Link href="/dashboard/locacao/inquilinos">
      <Button>
        Voltar
      </Button>
    </Link>
  </div>
}