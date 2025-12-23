import { Loader2Icon } from "lucide-react"
import { IconLoader, IconLoader3, IconLoaderQuarter } from "@tabler/icons-react"

import { cn } from "@/lib/utils"

function Spinner({ className, ...props }: React.ComponentProps<"svg">) {
  return (
    <IconLoader
      role="status"
      aria-label="Loading"
      className={cn("size-4 animate-spin", className)}
      {...props}
    />
  )
}

export { Spinner }
