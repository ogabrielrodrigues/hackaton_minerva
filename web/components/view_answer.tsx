'use client'

import { useState } from 'react'
import { ViewAnswerModal } from './modal/view_answer_modal'

interface Modal {
  open: boolean
  feedback_id: string
}

interface Props {
  feedback_id: string
  answer_id: string
  token: string
}

export function ViewAnswer({ feedback_id, answer_id, token }: Props) {
  const [view_answer_modal, setViewAnswerModal] = useState<Modal>({
    open: false,
    feedback_id,
  })

  function toggleModal() {
    setViewAnswerModal({
      open: !view_answer_modal.open,
      feedback_id,
    })
  }

  return (
    <>
      <button
        onClick={toggleModal}
        className="px-2 py-1 bg-primary text-zinc-50 rounded cursor-pointer"
      >
        Ver resposta
      </button>
      <ViewAnswerModal
        open={view_answer_modal.open}
        answer_id={answer_id}
        token={token}
        toggleModal={toggleModal}
      />
    </>
  )
}
