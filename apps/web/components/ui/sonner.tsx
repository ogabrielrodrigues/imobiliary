"use client"

import { useTheme } from "next-themes"
import { Toaster as Sonner, ToasterProps } from "sonner"

const Toaster = ({ ...props }: ToasterProps) => {
  const { theme = "system" } = useTheme()

  return (
    <Sonner
      theme={theme as ToasterProps["theme"]}
      className="toaster group"
      toastOptions={{
        classNames: {
          title: "!text-primary !font-bold",
          toast: "!bg-popover !border-border",
          description: "!text-muted-foreground",
          icon: "!text-primary !mr-2"
        }
      }}
      {...props}
    />
  )
}

export { Toaster }

