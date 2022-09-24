# uniar

`uniar` is database and management your scene collections CLI tool for [UNI'S ON AIR](https://keyahina-unisonair.com/).

## Usage

```
$ uniar list music | head
+-----------------------------------+-----------------------------------+--------+--------+----------+--------+
|               LIVE                |               MUSIC               |  TYPE  | LENGTH |  BONUS   | MASTER |
+-----------------------------------+-----------------------------------+--------+--------+----------+--------+
| 欅坂46                            | Student Dance                     | Blue   |    128 | {0 true} |     21 |
| 夏の全国アリーナツアー2018        |                                   |        |        |          |        |
| 欅坂46                            | AM1:27                            | Yellow |    127 | {0 true} |     21 |
| 夏の全国アリーナツアー2018        |                                   |        |        |          |        |
| 欅坂46                            | エキセントリック                  | Purple |    197 | {0 true} |     23 |
| 夏の全国アリーナツアー2018        |                                   |        |        |          |        |
| 欅坂46                            | ガラスを割れ！                    | Red    |    132 | {1 true} |     24 |

```

## Commands

### Setup

Setup your member status, office bonus status, scene card collections.

```
$ uniar setup
== メンバーステータスセットアップ ==
石森虹花
絆ランク (現在値:100) [100]:
```

### List

Show database data

- group
- member
- live
- music
- scene

### Regist

- live
- music
- photograph
- scene

### Update

TBD
