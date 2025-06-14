"""Add inverted index tables

Revision ID: e0e2349d109a
Revises: 8f2c33eb4bcc
Create Date: 2025-06-04 14:52:41.305968

"""
from typing import Sequence, Union

from alembic import op
import sqlalchemy as sa


# revision identifiers, used by Alembic.
revision: str = 'e0e2349d109a'
down_revision: Union[str, None] = '8f2c33eb4bcc'
branch_labels: Union[str, Sequence[str], None] = None
depends_on: Union[str, Sequence[str], None] = None


def upgrade() -> None:
    """Upgrade schema."""
    # ### commands auto generated by Alembic - please adjust! ###
    op.create_table('dictionary',
    sa.Column('word_id', sa.Integer(), autoincrement=True, nullable=False),
    sa.Column('word', sa.String(), nullable=False),
    sa.PrimaryKeyConstraint('word_id')
    )
    op.create_index(op.f('ix_dictionary_word'), 'dictionary', ['word'], unique=True)
    op.create_table('freqtable',
    sa.Column('word_id', sa.Integer(), nullable=False),
    sa.Column('repo_id', sa.Integer(), nullable=False),
    sa.Column('freq', sa.Integer(), nullable=False),
    sa.PrimaryKeyConstraint('word_id', 'repo_id')
    )
    op.create_table('invertedindex',
    sa.Column('word_id', sa.Integer(), nullable=False),
    sa.Column('dfi', sa.Integer(), nullable=False),
    sa.PrimaryKeyConstraint('word_id')
    )
    # ### end Alembic commands ###


def downgrade() -> None:
    """Downgrade schema."""
    # ### commands auto generated by Alembic - please adjust! ###
    op.drop_table('invertedindex')
    op.drop_table('freqtable')
    op.drop_index(op.f('ix_dictionary_word'), table_name='dictionary')
    op.drop_table('dictionary')
    # ### end Alembic commands ###
