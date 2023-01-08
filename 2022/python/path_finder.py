from __future__ import annotations

from random import randint


class PathFinder:
    def __init__(self, map):
        self.map = map

    def find(self):
        best_path = None
        paths = [self.map.get_start()]
        lengths = {}
        i = 0
        while paths:
            i += 1
            path = paths.pop()
            if randint(1, 5000) == 1:
                print(i, len(paths), len(path))

            if self.map.is_done(path):
                if best_path is None or self.map.evaluate_best_cost(path) < self.map.evaluate_best_cost(best_path):
                    best_path = path

                continue

            if best_path is not None and self.map.evaluate_best_cost(
                path
            ) > self.map.evaluate_best_cost(best_path):
                continue

            for next_path in self.map.available_paths(path):
                state = next_path.current_state
                if state in lengths and lengths[state] <= len(next_path):
                    continue

                paths.append(next_path)
                lengths[state] = len(state)

            paths.sort(key=lambda p: self.map.evaluate_best_cost(p), reverse=True)

        return best_path
