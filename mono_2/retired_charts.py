# AUDITSCAN FIXTURE -- dead chart helper using pillow. PLAN-C must keep
# pillow not_affected WITH downgrade proof naming this function (pillow is
# otherwise unused, so import-level analysis already says not_affected --
# the point is the proof trail, not the verdict).
import PIL.Image


def _retired_render_2019():
    return PIL.Image.open("legacy.png")
