'use client'

import { signAction } from '@/actions/sign'
import { redirect, useRouter } from 'next/navigation'
import { useFormStatus } from 'react-dom'
import { useFormState } from 'react-dom'

const sign_status = {
  success: false,
}

function Submit() {
  const { pending } = useFormStatus()

  return (
    <button
      type="submit"
      aria-disabled={pending}
      className="h-10 bg-primary text-zinc-50 rounded cursor-pointer"
    >
      Cadastrar
    </button>
  )
}

export function SignForm() {
  const [, sign_action] = useFormState(signAction, sign_status)
  const navigate = useRouter()

  return (
    <form
      action={async fd => {
        await sign_action(fd)

        navigate.push('/')
      }}
      className="flex flex-col gap-2"
    >
      <div className="flex flex-col gap-1">
        <label htmlFor="name" className="font-medium">
          Nome completo
        </label>
        <input
          id="name"
          name="name"
          type="text"
          className="h-10 border border-zinc-200 rounded pl-2 outline-none focus:ring-2 focus:ring-primary transition-all"
          placeholder="Digite seu nome completo"
        />
      </div>
      <div className="flex flex-col gap-1">
        <label htmlFor="email" className="font-medium">
          E-mail
        </label>
        <input
          id="email"
          name="email"
          type="text"
          className="h-10 border border-zinc-200 rounded pl-2 outline-none focus:ring-2 focus:ring-primary transition-all"
          placeholder="Digite seu e-mail"
        />
      </div>
      <div className="flex flex-col gap-1">
        <label htmlFor="unit" className="font-medium">
          Unidade
        </label>
        <input
          id="unit"
          name="unit"
          type="text"
          className="h-10 border border-zinc-200 rounded pl-2 outline-none focus:ring-2 focus:ring-primary transition-all"
          placeholder="Digite sua unidade"
        />
      </div>
      <div className="flex flex-col gap-1">
        <label htmlFor="sector" className="font-medium">
          Setor
        </label>
        <input
          id="sector"
          name="sector"
          type="text"
          className="h-10 border border-zinc-200 rounded pl-2 outline-none focus:ring-2 focus:ring-primary transition-all"
          placeholder="Digite seu setor"
        />
      </div>
      <div className="flex flex-col gap-1">
        <label htmlFor="password" className="font-medium">
          Senha
        </label>
        <input
          id="password"
          name="password"
          type="password"
          className="h-10 border border-zinc-200 rounded pl-2 outline-none focus:ring-2 focus:ring-primary transition-all"
          placeholder="Digite sua senha"
        />
      </div>
      <Submit />
    </form>
  )
}
