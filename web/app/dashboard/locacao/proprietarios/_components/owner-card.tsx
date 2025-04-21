import { Button } from "@/components/ui/button"
import { Card, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"
import { Owner } from "@/types/owner"
import { ArrowUpRight } from "lucide-react"
import Link from "next/link"

type OwnerCardProps = {
  owner: Owner
}

export function OwnerCard({ owner }: OwnerCardProps) {
  return (
    <Card className="bg-zinc-900/20 backdrop-blur-2xl relative z-20 overflow-hidden">
      <CardHeader className="flex flex-col h-20">
        <CardTitle>{owner.fullname}</CardTitle>
        <CardDescription>{owner.address.mini_address}</CardDescription>
      </CardHeader>
      <CardFooter className="flex justify-end">
        <div className="flex justify-end">
          <Link href={`/dashboard/locacao/proprietarios/${owner.id}`}>
            <Button variant="outline">
              <ArrowUpRight className="size-4 mr-1" />
              <p className="hidden sm:block">Detalhes</p>
            </Button>
          </Link>
        </div>
      </CardFooter>
    </Card>
  )
}