'use client'

import { updateAvatar } from "@/actions/avatar";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { Form, FormControl, FormField, FormItem } from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Tooltip, TooltipContent, TooltipTrigger } from "@/components/ui/tooltip";
import { env } from "@/lib/env";
import { User } from "@/types/user";
import { zodResolver } from "@hookform/resolvers/zod";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { toast } from "sonner";
import { z } from "zod";

type AvatarFormProps = {
  user: User | undefined
}

const avatar_schema = z.object({
  avatar: z.instanceof(File)
    .refine(file => file.size > 0 && file.size < 1024 * 1024 * 3),
})

export function AvatarForm({ user }: AvatarFormProps) {
  const [avatar, setAvatar] = useState<string | undefined>(user?.avatar)
  const form = useForm<z.infer<typeof avatar_schema>>({
    resolver: zodResolver(avatar_schema),
  })

  async function onSubmit(values: z.infer<typeof avatar_schema>) {
    const formData = new FormData()
    formData.append("avatar", values.avatar)
    const { status, url } = await updateAvatar(formData)

    if (status === 401) {
      toast.error("Não autorizado")
      return
    }

    if (status === 200) {
      toast.success("Avatar alterado com sucesso")
      setAvatar(url)
    }

    form.reset()
  }

  return (
    <Form {...form}>
      <form>
        <label htmlFor="avatar" className="cursor-pointer">
          <Tooltip>
            <TooltipTrigger asChild>
              <Avatar className="w-20 h-20 text-2xl font-semibold hover:opacity-80 transition-opacity">
                {
                  avatar ?
                    <AvatarImage src={`${env.SERVER_ADDR}/users/avatar/${avatar}`} className="object-cover" /> :
                    <AvatarFallback className="bg-sidebar-primary">{user?.fullname?.charAt(0)}</AvatarFallback>
                }
              </Avatar>
            </TooltipTrigger>
            <TooltipContent side="bottom">
              <p>Alterar avatar</p>
            </TooltipContent>
          </Tooltip>
          <FormField
            control={form.control}
            name="avatar"
            render={({ field: { onChange, value, ...fieldProps } }) => (
              <FormItem>
                <FormControl>
                  <Input
                    type="file"
                    id="avatar"
                    accept="image/*"
                    className="hidden"
                    {...fieldProps}
                    onChange={e => {
                      onChange(e.target.files?.[0])
                      form.handleSubmit(onSubmit)()
                    }}
                  />
                </FormControl>
              </FormItem>
            )}
          />
        </label>
      </form>
    </Form>
  )
}
