
# uniar

`uniar` is database and management your scene collections CLI tool for [UNI'S ON AIR](https://keyahina-unisonair.com/).

## Usage		

```
uniar is UNI'S ON AIR music and scene cards database and manage your scene cards collection tool.

Usage:
  uniar [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  doc         Generate uniar document
  help        Help about any command
  list        Show data
  regist      Regist data
  server      Server
  setup       Setup uniar

Flags:
      --config string   config file (default is $HOME/.uniar.yaml)
  -h, --help            help for uniar
  -v, --version         version for uniar

Use "uniar [command] --help" for more information about a command.

```

## Install

```
$ brew tap litencatt/tap
$ brew install litencatt/tap/uniar
```

## Commands

### List

```
$ uniar list
uniar is UNI'S ON AIR music and scene cards database and manage your scene cards collection tool.

Usage:
  uniar [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  doc         Generate uniar document
  help        Help about any command
  list        Show data
  regist      Regist data
  server      Server
  setup       Setup uniar

Flags:
      --config string   config file (default is $HOME/.uniar.yaml)
  -h, --help            help for uniar
  -v, --version         version for uniar

Use "uniar [command] --help" for more information about a command.

```

### Usage

```
$ uniar list scene -h
Show scene card list

Usage:
  uniar list scene [flags]

Aliases:
  scene, s

Flags:
  -c, --color string            Color filter(e.g. -c Red or -c r)
  -d, --detail                  Show detail
  -f, --full-name               Show pohtograph full name
      --have                    Show only scenes you have
  -h, --help                    help for scene
  -i, --ignore-columns string   Ignore columns to display(VoDa50,DaPe50,...)
  -m, --member string           Member filter(e.g. -m 加藤史帆)
  -n, --not-have                Show only scenes you NOT have
  -p, --photograph string       Photograph filter(e.g. -p JOYFULLOVE)
  -s, --sort string             Sort target rank.(all35, voda50, ...)

Global Flags:
      --config string   config file (default is $HOME/.uniar.yaml)

```

```
$ uniar list scene -f | head
+--------+------------------------------------+---------+------------+--------+-------+-------+--------+--------+--------+------+------+------+-------+
| COLOR  |             PHOTOGRAPH             | SSRPLUS |   MEMBER   | EXPECT | TOTAL | ALL35 | VODA50 | DAPE50 | VOPE50 | VO85 | DA85 | PE85 | HAVE  |
+--------+------------------------------------+---------+------------+--------+-------+-------+--------+--------+--------+------+------+------+-------+
| Blue   | 風に吹かれても                     | true    | 上村莉菜   |   3.68 | 12931 |     1 |      2 |      5 |      2 |    1 |    3 |    5 | false |
| Red    | アザトカワイイ                     | true    | 佐々木美玲 |   3.72 | 13324 |     2 |     28 |      1 |      1 |   24 |   30 |    1 | false |
| Purple | Buddies                            | true    | 山﨑天     |   3.72 | 13323 |     3 |      1 |      2 |     28 |   23 |    1 |   23 | false |
| Blue   | 青春の馬                           | true    | 濱岸ひより |   3.72 | 13077 |     4 |     34 |      4 |      4 |   38 |   47 |    3 | false |
| Purple | なぜ　恋をして来なかったんだろう？ | true    | 藤吉夏鈴   |   3.72 | 13076 |     5 |     36 |      3 |      3 |   39 |   48 |    2 | false |
| Yellow | 誰がその鐘を鳴らすのか？           | true    | 山﨑天     |   3.68 | 13076 |     6 |      5 |     25 |      6 |    2 |   22 |   18 | false |
| Green  | こんなに好きになっちゃっていいの？ | true    | 加藤史帆   |   3.68 | 13067 |     7 |      4 |     26 |      7 |    3 |   21 |   20 | false |

```

```
$ uniar list scene -f -c Blue
+-------+------------------------------------+---------+------------+--------+-------+-------+--------+--------+--------+------+------+------+-------+
| COLOR |             PHOTOGRAPH             | SSRPLUS |   MEMBER   | EXPECT | TOTAL | ALL35 | VODA50 | DAPE50 | VOPE50 | VO85 | DA85 | PE85 | HAVE  |
+-------+------------------------------------+---------+------------+--------+-------+-------+--------+--------+--------+------+------+------+-------+
| Blue  | 風に吹かれても                     | true    | 上村莉菜   |   3.68 | 12931 |     1 |      1 |      2 |      1 |    1 |    2 |    3 | false |
| Blue  | 青春の馬                           | true    | 濱岸ひより |   3.72 | 13077 |     2 |      6 |      1 |      2 |    4 |    7 |    1 | false |
| Blue  | Nobody’s fault                    | true    | 渡邉理佐   |   3.68 | 13065 |     3 |      2 |      3 |      6 |    5 |    1 |    8 | false |
| Blue  | アンビバレント                     | true    | 田村保乃   |   3.68 | 12976 |     4 |      5 |      4 |      4 |    6 |    4 |    2 | false |
| Blue  | ドレミソラシド                     | true    | 森本茉莉   |   3.68 | 12974 |     5 |      3 |      6 |      3 |    2 |    5 |    4 | false |
| Blue  | ガラスを割れ！                     | true    | 土生瑞穂   |   3.68 | 12819 |     6 |      4 |      5 |      5 |    3 |    3 |    7 | false |
| Blue  | UNI'S ON AIR 3rd ANNIVERSARY       | false   | 潮紗理菜   |   3.68 | 11435 |     7 |      7 |     14 |     12 |    8 |   10 |   29 | false |

```

```
$ uniar list scene -f -m 加藤史帆
+--------+------------------------------------+---------+----------+--------+-------+-------+--------+--------+--------+------+------+------+-------+
| COLOR  |             PHOTOGRAPH             | SSRPLUS |  MEMBER  | EXPECT | TOTAL | ALL35 | VODA50 | DAPE50 | VOPE50 | VO85 | DA85 | PE85 | HAVE  |
+--------+------------------------------------+---------+----------+--------+-------+-------+--------+--------+--------+------+------+------+-------+
| Green  | こんなに好きになっちゃっていいの？ | true    | 加藤史帆 |   3.68 | 13067 |     1 |      1 |      1 |      1 |    1 |    1 |    1 | false |
| Yellow | 君しか勝たん                       | false   | 加藤史帆 |   3.68 | 11083 |     2 |      2 |      2 |      4 |    4 |    2 |    5 | false |
| Green  | こんなに好きになっちゃっていいの？ | false   | 加藤史帆 |   3.68 | 10817 |     3 |      3 |      5 |      2 |    2 |    5 |    4 | false |
| Blue   | NEW YEAR'23                        | false   | 加藤史帆 |   3.49 | 11269 |     4 |      4 |      4 |      3 |    3 |    4 |    3 | false |
| Purple | 清夏彩る涼の装                     | false   | 加藤史帆 |    3.4 | 11357 |     5 |      5 |      3 |      5 |    5 |    3 |    2 | false |
| Red    | ハッピーオーラ                     | true    | 加藤史帆 |   2.46 | 12932 |     6 |      6 |      6 |      6 |    6 |    7 |    6 | false |
| Blue   | キュン                             | true    | 加藤史帆 |   2.46 | 12560 |     7 |      7 |      7 |      7 |    7 |    8 |    9 | false |

```

```
$ uniar list scene -f -p キュン
+--------+------------+---------+------------+--------+-------+-------+--------+--------+--------+------+------+------+-------+
| COLOR  | PHOTOGRAPH | SSRPLUS |   MEMBER   | EXPECT | TOTAL | ALL35 | VODA50 | DAPE50 | VOPE50 | VO85 | DA85 | PE85 | HAVE  |
+--------+------------+---------+------------+--------+-------+-------+--------+--------+--------+------+------+------+-------+
| Green  | キュン     | true    | 小坂菜緒   |   3.68 | 12582 |     1 |      1 |      1 |      1 |    1 |    1 |    1 | false |
| Green  | キュン     | false   | 小坂菜緒   |   3.68 | 10332 |     2 |      2 |      2 |      2 |    2 |    2 |    2 | false |
| Red    | キュン     | true    | 松田好花   |   2.75 | 12902 |     3 |      3 |      3 |      3 |    3 |    3 |    3 | false |
| Purple | キュン     | true    | 上村ひなの |   2.52 | 12768 |     4 |      4 |      4 |      5 |    5 |    4 |    4 | false |
| Green  | キュン     | true    | 佐々木美玲 |   2.52 | 12599 |     5 |      5 |      5 |      4 |    4 |    6 |    5 | false |
| Blue   | キュン     | true    | 加藤史帆   |   2.46 | 12560 |     6 |      6 |      6 |      6 |    6 |    5 |    6 | false |
| Purple | キュン     | true    | 潮紗理菜   |   2.32 | 12647 |     7 |      7 |      7 |      7 |    7 |    7 |    7 | false |

```

```
$ uniar list scene -f -c Blue -m 加藤史帆 -p キュン
+-------+------------+---------+----------+--------+-------+-------+--------+--------+--------+------+------+------+-------+
| COLOR | PHOTOGRAPH | SSRPLUS |  MEMBER  | EXPECT | TOTAL | ALL35 | VODA50 | DAPE50 | VOPE50 | VO85 | DA85 | PE85 | HAVE  |
+-------+------------+---------+----------+--------+-------+-------+--------+--------+--------+------+------+------+-------+
| Blue  | キュン     | true    | 加藤史帆 |   2.46 | 12560 |     1 |      1 |      1 |      1 |    1 |    1 |    1 | false |
| Blue  | キュン     | false   | 加藤史帆 |   2.46 | 10310 |     2 |      2 |      2 |      2 |    2 |    2 |    2 | false |
+-------+------------+---------+----------+--------+-------+-------+--------+--------+--------+------+------+------+-------+

```

```
$ uniar list scene --ignore-columns All35,VoDa50,DaPe50,VoPe50 | head
+--------+--------------------------------+---------+------------+--------+-------+------+------+------+-------+
| COLOR  |           PHOTOGRAPH           | SSRPLUS |   MEMBER   | EXPECT | TOTAL | VO85 | DA85 | PE85 | HAVE  |
+--------+--------------------------------+---------+------------+--------+-------+------+------+------+-------+
| Blue   | 風吹か                         | true    | 上村莉菜   |   3.68 | 12931 |    1 |    3 |    5 | false |
| Red    | アザカワ                       | true    | 佐々木美玲 |   3.72 | 13324 |   24 |   30 |    1 | false |
| Purple | Buddies                        | true    | 山﨑天     |   3.72 | 13323 |   23 |    1 |   23 | false |
| Blue   | 馬                             | true    | 濱岸ひより |   3.72 | 13077 |   38 |   47 |    3 | false |
| Purple | なぜ恋                         | true    | 藤吉夏鈴   |   3.72 | 13076 |   39 |   48 |    2 | false |
| Yellow | 誰鐘                           | true    | 山﨑天     |   3.68 | 13076 |    2 |   22 |   18 | false |
| Green  | こん好き                       | true    | 加藤史帆   |   3.68 | 13067 |    3 |   21 |   20 | false |

```

```
$ uniar list scene -d --ignore-columns All35,VoDa50,DaPe50,VoPe50 | head
+--------+--------------------------------+---------+------------+--------+-------+------+------+------+------+------+------+-------+
| COLOR  |           PHOTOGRAPH           | SSRPLUS |   MEMBER   | EXPECT | TOTAL | VO85 | DA85 | PE85 |  VO  |  DA  |  PE  | HAVE  |
+--------+--------------------------------+---------+------------+--------+-------+------+------+------+------+------+------+-------+
| Blue   | 風吹か                         | true    | 上村莉菜   |   3.68 | 12931 |    1 |    3 |    5 | 4776 | 3883 | 3842 | false |
| Red    | アザカワ                       | true    | 佐々木美玲 |   3.72 | 13324 |   24 |   30 |    1 | 3023 | 3036 | 6835 | false |
| Purple | Buddies                        | true    | 山﨑天     |   3.72 | 13323 |   23 |    1 |   23 | 3033 | 6795 | 3065 | false |
| Blue   | 馬                             | true    | 濱岸ひより |   3.72 | 13077 |   38 |   47 |    3 | 2989 | 3015 | 6643 | false |
| Purple | なぜ恋                         | true    | 藤吉夏鈴   |   3.72 | 13076 |   39 |   48 |    2 | 2971 | 2974 | 6701 | false |
| Yellow | 誰鐘                           | true    | 山﨑天     |   3.68 | 13076 |    2 |   22 |   18 | 5092 | 3756 | 3798 | false |
| Green  | こん好き                       | true    | 加藤史帆   |   3.68 | 13067 |    3 |   21 |   20 | 5102 | 3780 | 3755 | false |

```
