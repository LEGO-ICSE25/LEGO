from .base import BaseFormatter


class Simple(BaseFormatter):
    def __init__(self, cg_generator):
        self.cg_generator = cg_generator

    def generate(self):
        output = self.cg_generator.output()
        output_cg = {}
        for node in output:
            output_cg[node] = list(output[node])
            #print(node, output_cg[node])
        return output_cg
