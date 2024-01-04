'use client'

import { filterAction } from '@/actions/filter'
import { useRef } from 'react'
import { useFormState, useFormStatus } from 'react-dom'

const filter = {
  filtered: false,
}

interface Props {
  filtered: boolean
}

function SubmitButton({ filtered }: Props) {
  const { pending } = useFormStatus()

  return (
    <button
      type="submit"
      aria-disabled={pending}
      className="h-10 w-full bg-primary text-zinc-50 rounded cursor-pointer"
    >
      {filtered ? 'Reset' : 'Filtrar'}
    </button>
  )
}

export function FilterForm() {
  const [filters, filter_action] = useFormState(filterAction, filter)
  const ref = useRef<HTMLFormElement>(null)

  return (
    <form
      ref={ref}
      action={async fd => {
        await filter_action(fd)
        ref.current?.reset()
      }}
      className="w-full flex flex-col gap-2 "
    >
      <label htmlFor="sector" className="font-medium">
        Filtrar por Setor
      </label>
      <div className="flex gap-2">
        <input
          id="sector"
          name="sector"
          type="text"
          autoComplete="off"
          className="h-10 border border-zinc-200 rounded pl-2 outline-none focus:ring-2 focus:ring-primary transition-all"
          placeholder="Digite o setor"
        />
        <SubmitButton filtered={filters.filtered} />
      </div>
    </form>
  )
}
