'use client'

import { X } from 'lucide-react'
import { Answer } from '../answer'

interface Props {
  open: boolean
  answer_id: string
  token: string
  toggleModal(): void
}

export function ViewAnswerModal({
  open,
  answer_id,
  token,
  toggleModal,
}: Props) {
  return (
    <div
      data-open={open}
      className="hidden data-[open=true]:flex w-screen h-screen absolute top-0 left-0 bg-zinc-900/50 z-10"
    >
      <section
        data-open={open}
        className="hidden data-[open=true]:flex flex-col w-[90%] h-fit p-4 absolute top-1/2 left-1/2 -translate-y-1/2 -translate-x-1/2  rounded bg-zinc-100 border border-zinc-200 z-50 shadow-lg"
      >
        <div className="w-full h-fit flex items-center justify-end">
          <button
            onClick={toggleModal}
            className="p-2 text-primary font-medium hover:bg-zinc-300/40 rounded transition-colors"
          >
            <X />
          </button>
        </div>
        <h2 className="text-4xl font-bold text-primary">Resposta</h2>
        <Answer answer_id={answer_id} token={token} />
      </section>
    </div>
  )
}
