import { Button } from "@/components/ui/button";
import { HousePlus } from "lucide-react";
import Link from "next/link";

type AuthHeaderProps = {
  title: string
  description: string
  url: string
  urlText: string
}

export function AuthHeader({ title, description, url, urlText }: AuthHeaderProps) {
  return (
    <div className="flex flex-col gap-6">
      <div className="flex flex-col items-center gap-2">
        <div className="flex flex-col items-center gap-2 font-medium">
          <Link href="/" className="flex size-8 items-center justify-center rounded-md">
            <Button variant="ghost" size="icon">
              <HousePlus className="size-8" />
            </Button>
          </Link>
          <span className="sr-only">Imobiliary</span>
        </div>
        <h1 className="text-xl font-bold text-center">{title}</h1>
        <div className="text-center text-sm">
          {description}{" "}
          <a href={url} id="auth-option" className="underline underline-offset-4">
            {urlText}
          </a>
        </div>
      </div>
    </div>
  )
}
