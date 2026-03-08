from enum import Enum

class OrderedEnum(Enum):
    def __ge__(self, other):
        if self.__class__ is other.__class__:
            return self.value >= other.value
        return NotImplemented
    def __gt__(self, other):
        if self.__class__ is other.__class__:
            return self.value > other.value
        return NotImplemented
    def __le__(self, other):
        if self.__class__ is other.__class__:
            return self.value <= other.value
        return NotImplemented
    def __lt__(self, other):
        if self.__class__ is other.__class__:
            return self.value < other.value
        return NotImplemented

class Precedence(OrderedEnum):
    LOWEST = 1
    ADD = 2
    MULT = 3

def tok_precedence(token):
    if token in ['*', '/']:
        return Precedence.MULT
    elif token in ['+', '-']:
        return Precedence.ADD
    else:
        return Precedence.LOWEST

class Expr:
    def __init__(self, left, right, op):
        self.left = left
        self.right = right
        self.op = op

    def __str__(self):
        return f'({self.left} {self.op} {self.right})'

class PrattParser:
    def __init__(self, tokens=[]):
        self.tokens = tokens
        self.idx = 0

    @property
    def curr_tok(self):
        return self.tokens[self.idx]

    @property
    def peek_tok(self):
        return self.tokens[self.idx+1]

    def next_tok(self):
        self.idx += 1

    def parsePrefix(self):
        if self.curr_tok == '(':
            self.next_tok()

            expr = self.parseExpr(Precedence.LOWEST)

            self.next_tok()

            return expr
        else:
            tok = self.curr_tok
            self.next_tok()
            return tok

    def parseExpr(self, prec):
        left = self.parsePrefix()

        while self.idx < len(self.tokens) and prec < tok_precedence(self.curr_tok):
            op = self.curr_tok
            op_prec = tok_precedence(op)

            self.next_tok()

            right = self.parseExpr(op_prec)

            left = Expr(left, right, op)

        return left

    @staticmethod
    def parse(toks):
        return PrattParser(toks).parseExpr(Precedence.LOWEST)

toks = '1 + 2 + 3 + ( 6 + 7 ) * 4 + 5'.split()

expr = PrattParser.parse(toks)

print(expr)
