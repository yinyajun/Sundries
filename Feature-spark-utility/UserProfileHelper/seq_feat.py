#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2019/9/20 18:10
# @Author  : Yajun Yin
# @Note    :


class SeqFeature(object):
    def __init__(self, threshold=0, truncate=None, has_duration=None):
        self._hist = []
        self._weights = []
        self.th = threshold
        self.truncate = truncate
        self.has_duration = has_duration
        if self.has_duration:
            self._durations = []

    def add(self, item, score=None, duration=None):
        if not score or score > self.th:
            self._hist.append(item)
            s = score if score else 1
            assert s >= 0
            self._weights.append(s)
            if self.has_duration:
                assert duration is not None, "duration is None when has_duration set true."
                self._durations.append(duration)

    @staticmethod
    def _normalized_weight(weights):
        s = float(sum(weights))
        return [i / s for i in weights]

    def _deduplicate(self):
        # Remove relative order information in history sequence.
        hist = {}
        for i in range(len(self._hist)):
            item = self._hist[i]
            cnt = self._weights[i]
            hist[item] = hist.get(item, 0) + cnt
        ret = sorted(hist.items(), key=lambda p: p[1], reverse=True)
        if self.truncate:
            ret = ret[:self.truncate]
        history, weights = zip(*ret)
        return history, weights

    def _treat_ordered_seq(self, decay_strategy=None):
        history = self._hist
        weights = self._weights
        if self.has_duration:
            durations = self._durations
            # sort by duration
            data = zip(history, weights, durations)
            data = sorted(data, key=lambda p: p[2])
            if self.truncate:
                data = data[:self.truncate]
            for i in range(len(data)):
                dur = durations[i]
                decay = decay_strategy(dur) if decay_strategy else 1
                weights[i] = weights[i] * decay
            return history, weights
        return history, weights

    def output(self, keep_order=False, normalize=True, decay_strategy=None):
        if not keep_order:
            history, weights = self._deduplicate()
            if normalize:
                weights = self._normalized_weight(weights)
            return history, weights
        history, weights = self._treat_ordered_seq(decay_strategy=decay_strategy)
        if normalize:
            weights = self._normalized_weight(weights)
        return history, weights


def cooling_strategy(H, alpha, T0, tao):
    from math import exp
    # adopt Newton's law of cooling
    # T = H + (T0 - H) exp ( - alpha * ( t - t0 ) )
    return H + (T0 - H) * exp(-alpha * tao)


def strategy(tao):
    return cooling_strategy(0.5, 0.05, 1, tao)


if __name__ == '__main__':
    f = SeqFeature(truncate=10)
    f.add('a', 2)
    f.add('b', 1)
    f.add('a', 3)
    f.add('d', 1)
    f.add('c', 6)
    print(f.output())