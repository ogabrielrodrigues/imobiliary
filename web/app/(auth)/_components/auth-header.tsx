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
            <Button variant="ghost" size="sm" className="flex items-center gap-2">
              <HousePlus className="size-5" />
              <span className="sr-only">Imobiliary</span>
              <h1 className="text-2xl font-bold select-none">Imobiliary</h1>
            </Button>
          </Link>
        </div>
        <div className="text-center text-sm mt-4 text-muted-foreground">
          {description}{" "}
          <a href={url} id="auth-option" className="underline underline-offset-4 text-primary">
            {urlText}
          </a>
        </div>
      </div>
    </div>
  )
}
