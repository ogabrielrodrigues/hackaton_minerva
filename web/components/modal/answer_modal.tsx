'use client'

import { X } from 'lucide-react'
import { AnswerForm } from '../forms/answer_form'

interface Props {
  open: boolean
  feedback_id: string
  toggleModal(): void
}

export function AnswerModal({ open, feedback_id, toggleModal }: Props) {
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
        <h2 className="text-4xl font-bold text-primary">Responder</h2>

        <AnswerForm feedback_id={feedback_id} toggleModal={toggleModal} />
      </section>
    </div>
  )
}
