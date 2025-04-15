'use client'

import { updateAvatar } from "@/actions/avatar";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { Button } from "@/components/ui/button";
import { Form, FormControl, FormField, FormItem } from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Tooltip, TooltipContent, TooltipTrigger } from "@/components/ui/tooltip";
import { User } from "@/types/user";
import { zodResolver } from "@hookform/resolvers/zod";
import { useRouter } from "next/navigation";
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
  const router = useRouter()
  const form = useForm<z.infer<typeof avatar_schema>>({
    resolver: zodResolver(avatar_schema),
  })

  async function onSubmit(values: z.infer<typeof avatar_schema>) {
    const formData = new FormData()
    formData.append("avatar", values.avatar)
    const status = await updateAvatar(formData)

    switch (status) {
      case 200:
        toast.success("Avatar atualizado com sucesso", {
          description: "Atualize sua conta para ver as mudanças",
          duration: 5000,
          action: <Button variant="outline" onClick={() => window.location.reload()}>Atualizar</Button>,
        })
        setTimeout(() => {
          window.location.reload()
        }, 5000)
        break
      case 400:
        toast.error("O arquivo deve ser uma imagem com tamanho máximo de 3MB")
        break
      case 500:
        toast.error("Erro ao atualizar o avatar")
        break
    }

    form.reset({ avatar: undefined })
  }

  return (
    <Form {...form}>
      <form>
        <label htmlFor="avatar" className="cursor-pointer">
          <Tooltip>
            <TooltipTrigger asChild>
              <Avatar className="w-20 h-20 text-2xl font-semibold hover:opacity-80 transition-opacity">
                {user?.avatar ? <AvatarImage src={user?.avatar} className="object-cover" />
                  : <AvatarFallback className="bg-sidebar-primary animate-pulse" />}
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
